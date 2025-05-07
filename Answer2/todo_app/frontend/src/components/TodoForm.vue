<template>
    <div v-if="formVisible" class="bg-white shadow-md rounded p-4 mb-4 border border-gray-200">
        <h2 class="text-xl font-bold mb-4">
            {{ form.id ? "แก้ไขรายการ" : "เพิ่มรายการใหม่" }}
        </h2>

        <form @submit.prevent="handleSubmit">
            <div class="mb-2">
                <label class="block text-sm font-medium text-gray-700">ชื่อเรื่อง</label>
                <input v-model.trim="form.title" type="text" class="input" required />
            </div>

            <div class="mb-2">
                <label class="block text-sm font-medium text-gray-700">รายละเอียด</label>
                <textarea v-model="form.description" class="input" rows="3"></textarea>
            </div>

            <div class="mb-2">
                <label class="block text-sm font-medium text-gray-700">ความสำคัญ</label>
                <select v-model.number="form.priority" class="input" required>
                    <option :value="1">🔴 สำคัญมาก (ต้องรีบส่งด่วน)</option>
                    <option :value="2">🟡 ระดับกลาง (พอมีเวลา)</option>
                    <option :value="3">🟢 สำคัญน้อย (มีเวลาส่ง)</option>
                </select>
            </div>

            <div class="mb-4">
                <label class="inline-flex items-center">
                    <input type="checkbox" v-model="form.is_completed" class="mr-2" />
                    ทำเสร็จแล้ว
                </label>
            </div>

            <div class="flex space-x-2">
                <button type="submit" class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600">
                    บันทึก
                </button>

                <button type="button" class="bg-gray-300 text-gray-800 px-4 py-2 rounded hover:bg-gray-400"
                    @click="handleCancel">
                    ยกเลิก
                </button>
            </div>
        </form>
    </div>
</template>

<script setup>
import { reactive, watch, ref } from "vue";
import { toRaw } from "vue";

// ตัวแปรที่ใช้ในการควบคุมการแสดงผลฟอร์ม
const formVisible = ref(true);

const props = defineProps({
    initialTodo: {
        type: Object,
        default: () => ({
            id: null,
            title: "",
            description: "",
            priority: 2,
            is_completed: false,
        }),
    },
});

const emit = defineEmits(["submit", "cancel"]);

const form = reactive({
    id: props.initialTodo.id,
    title: props.initialTodo.title,
    description: props.initialTodo.description,
    priority: props.initialTodo.priority,
    is_completed: props.initialTodo.is_completed,
});

// การตรวจสอบและการอัพเดตค่าของ form เมื่อ props มีการเปลี่ยนแปลง
watch(
    () => props.initialTodo,
    (newVal) => {
        Object.assign(form, newVal);
    }
);

function handleSubmit() {
    if (!form.title.trim()) {
        alert("กรุณาใส่ชื่อเรื่อง");
        return;
    }
    console.log("🚀 Submitting form:", toRaw(form));
    emit("submit", toRaw(form));
}

function handleCancel() {
    console.log("🚀 Canceling form");
    formVisible.value = false; // ซ่อนฟอร์มเมื่อกดยกเลิก
    Object.assign(form, props.initialTodo); // รีเซ็ตค่าฟอร์ม
    emit("cancel"); // ส่ง event cancel ไปยัง parent
}
</script>

<style scoped>
.input {
    width: 100%;
    padding: 0.5rem;
    border: 1px solid #ccc;
    border-radius: 0.375rem;
}
</style>
