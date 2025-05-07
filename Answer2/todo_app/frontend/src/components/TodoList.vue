<template>
    <div class="bg-white p-6 rounded-xl shadow-md border border-gray-200">
        <div class="flex justify-between items-center mb-4">
            <h2 class="text-2xl font-bold text-gray-800">📋 รายการทั้งหมด</h2>
            <button @click="startAdd"
                class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg transition duration-200">
                ➕ เพิ่มงานใหม่
            </button>
        </div>

        <!-- ตัวกรองความสำคัญ -->
        <div class="mb-6 flex items-center gap-4">
            <label class="text-gray-700 font-medium">กรองความสำคัญ:</label>
            <select v-model="filterPriority"
                class="border rounded px-3 py-1 text-sm focus:outline-none focus:ring focus:border-blue-400">
                <option value="">ทั้งหมด</option>
                <option value="1">🔴 สำคัญมาก</option>
                <option value="2">🟡 ระดับกลาง</option>
                <option value="3">🟢 สำคัญน้อย</option>
            </select>
        </div>

        <!-- ฟอร์มเพิ่ม/แก้ไข -->
        <TodoForm v-if="formVisible" :key="formKey" @submit="handleSubmit" :initialTodo="editingTodo" />

        <!-- ไม่มีรายการ -->
        <div v-if="visibleTodos.length === 0" class="text-gray-400 text-center py-6 italic">
            ยังไม่มีงานในระบบ
        </div>

        <!-- รายการ -->
        <transition-group name="fade" tag="div" class="space-y-4 mt-4">
            <TodoItem v-for="todo in visibleTodos" :key="todo.id" :todo="todo" @edit="startEdit" @delete="handleDelete"
                @toggle="toggleComplete" />
        </transition-group>
    </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted, watch } from "vue";
import axios from "axios";
import TodoItem from "./TodoItem.vue";
import TodoForm from "./TodoForm.vue";

const emit = defineEmits(["update-count"]);

const todos = ref([]);
const formVisible = ref(false);
const editingTodo = ref(null);
const filterPriority = ref("");


const visibleTodos = computed(() => {
    if (!Array.isArray(todos.value)) return [];
    return todos.value
        .filter((t) => !t.is_completed)
        .filter((t) =>
            filterPriority.value ? t.priority === Number(filterPriority.value) : true
        )
        .sort((a, b) => a.priority - b.priority);
});

async function fetchTodos() {
    const res = await axios.get("/todos");
    todos.value = Array.isArray(res.data) ? res.data : [];
}

function updateCount() {
    if (!Array.isArray(todos.value)) return;
    const total = todos.value.length;
    const completed = todos.value.filter((t) => t.is_completed).length;
    emit("update-count", { total, completed });
}

watch(
    [todos, filterPriority],
    () => {
        if (Array.isArray(todos.value)) updateCount();
    },
    { deep: true }
);

let isMounted = false;

onMounted(async () => {
    isMounted = true;
    await fetchTodos();
    if (isMounted) updateCount(); // ป้องกันเรียกตอน component ถูกถอดออก
});

onUnmounted(() => {
    isMounted = false;
});

function startAdd() {
    editingTodo.value = {
        title: "",
        description: "",
        priority: 2,
        is_completed: false,
    };
    console.log("🚀 Adding new todo");
    formVisible.value = true;
    formKey.value++; // เพิ่มการรีเฟรช key เพื่อให้คอมโพเนนต์ฟอร์มใหม่ถูกสร้าง
}


function startEdit(todo) {
    editingTodo.value = { ...todo };
    formVisible.value = true;
}

const formKey = ref(0);

function resetForm() {
    editingTodo.value = null;
    formKey.value++; // บังคับ Vue สร้าง component ใหม่
}

async function handleSubmit(todo) {
    try {
        if (todo.id) {
            await axios.put(`/todos/${todo.id}`, todo);
        } else {
            await axios.post("/todos", todo);
        }
        await fetchTodos();
        resetForm(); // รีเซ็ตฟอร์ม
        formVisible.value = false; // ซ่อนฟอร์มเมื่อบันทึกเสร็จ
    } catch (err) {
        console.error("บันทึกงานล้มเหลว", err);
    }
}


async function handleDelete(todo) {
    try {
        await axios.delete(`/todos/${todo.id}?confirm=true`);
        await fetchTodos(); // โหลดใหม่หลังลบ จะทำให้ visibleTodos ถูกอัปเดต
    } catch (err) {
        console.error("ลบงานผิดพลาด", err);
    }
}

async function toggleComplete(todo) {
    try {
        const updated = { ...todo, is_completed: !todo.is_completed };
        await axios.put(`/todos/${todo.id}`, updated);

        // ลบออกทันทีถ้าทำเสร็จ
        if (updated.is_completed) {
            todos.value = todos.value.filter((t) => t.id !== todo.id);
        } else {
            await fetchTodos();
        }
    } catch (err) {
        console.error("สลับสถานะล้มเหลว", err);
    }
}
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
    transition: all 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
    opacity: 0;
    transform: translateY(10px);
}
</style>
