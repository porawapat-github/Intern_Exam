<template>
    <div class="flex justify-between items-start p-4 rounded-lg border shadow-sm transition duration-200"
        :class="priorityClass">
        <div class="flex-1">
            <h3 class="text-lg font-semibold text-gray-800">
                {{ todo.title }}
                <span v-if="todo.is_completed" class="ml-2 text-green-600">✅</span>
            </h3>

            <p class="text-sm text-gray-600">{{ todo.description }}</p>

            <p class="text-xs text-gray-400 mt-1">
                แก้ไขล่าสุด: {{ formatDate(todo.updated_at) }}
            </p>
        </div>

        <div class="flex gap-2 items-center mt-2 sm:mt-0 sm:flex-row flex-col">
            <button class="text-green-600 hover:text-green-800 text-sm" @click="$emit('toggle', todo)">
                {{ todo.is_completed ? "↩️ ย้อนกลับ" : "✅ ทำสำเร็จ" }}
            </button>
            <button class="text-blue-600 hover:text-blue-800 text-sm" @click="$emit('edit', todo)">
                ✏️ แก้ไข
            </button>
            <button class="text-red-600 hover:text-red-800 text-sm" @click="$emit('delete', todo)">
                🗑️ ลบ
            </button>
        </div>
    </div>
</template>

<script setup>
import { computed } from "vue";
const props = defineProps({ todo: Object });

const priorityClass = computed(() => {
    switch (props.todo.priority) {
        case 1:
            return "border-red-500 bg-red-50";
        case 2:
            return "border-yellow-400 bg-yellow-50";
        case 3:
            return "border-green-500 bg-green-50";
        default:
            return "border-gray-300 bg-white";
    }
});

function formatDate(date) {
    if (!date) return "-";
    return new Date(date).toLocaleString("th-TH", {
        year: "numeric",
        month: "short",
        day: "numeric",
        hour: "2-digit",
        minute: "2-digit",
    });
}

</script>
