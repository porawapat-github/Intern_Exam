<template>
    <div class="max-w-4xl mx-auto p-6">
        <div class="mb-6 flex justify-between items-center">
            <router-link to="/" class="text-blue-600 hover:underline text-sm">
                ← กลับไปหน้าหลัก
            </router-link>
        </div>

        <TodoHistory />
    </div>
</template>

<script setup>
import TodoHistory from "../components/TodoHistory.vue";
</script>
