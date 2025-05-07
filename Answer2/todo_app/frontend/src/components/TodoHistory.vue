<template>
    <div class="bg-white p-6 rounded-xl shadow-md border border-gray-200">
        <div class="flex justify-between items-center mb-4">
            <h2 class="text-2xl font-bold text-gray-800">
                งานที่เสร็จแล้ว
            </h2>
        </div>

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

        <!-- แสดงรายการที่กรองแล้ว -->
        <transition-group name="fade" tag="div" class="space-y-4 mt-4">
            <TodoItem v-for="todo in filteredTodos" :key="todo.id" :todo="todo" @edit="editTodo" @delete="deleteTodo"
                @toggle="toggleComplete" />
        </transition-group>
    </div>
</template>

<script setup>
import { ref, onMounted, computed } from "vue";
import axios from "axios";
import TodoItem from "../components/TodoItem.vue";

const completedTodos = ref([]);
const filterPriority = ref("");

async function fetchCompleted() {
    try {
        const res = await axios.get("/todos");
        const data = Array.isArray(res.data) ? res.data : [];
        completedTodos.value = data.filter((t) => t.is_completed);
    } catch (err) {
        console.error("โหลดข้อมูลล้มเหลว", err);
        completedTodos.value = [];
    }
}

onMounted(fetchCompleted);

async function toggleComplete(todo) {
    try {
        const updated = { ...todo, is_completed: false };
        await axios.put(`/todos/${todo.id}`, updated);
        completedTodos.value = completedTodos.value.filter((t) => t.id !== todo.id);
    } catch (err) {
        console.error("อัปเดตสถานะล้มเหลว", err);
    }
}

async function deleteTodo(todo) {
    try {
        await axios.delete(`/todos/${todo.id}?confirm=true`);
        completedTodos.value = completedTodos.value.filter((t) => t.id !== todo.id);
    } catch (err) {
        console.error("ลบงานล้มเหลว", err);
    }
}

function editTodo(todo) {
    console.log("Edit Todo:", todo);
}

const filteredTodos = computed(() => {
    if (!filterPriority.value) return completedTodos.value;
    return completedTodos.value.filter(
        (t) => t.priority === Number(filterPriority.value)
    );
});
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
