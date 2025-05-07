<template>
  <div class="max-w-4xl mx-auto p-6 ">
    <div class="mb-6 flex justify-between items-center">
      <div>
        <h1 class="text-3xl font-bold text-gray-800">หน้าหลัก</h1>
        <p class="text-sm text-gray-500">
          งานทั้งหมด: {{ count.total }} | เสร็จแล้ว: {{ count.completed }}
        </p>
      </div>
    </div>

    <TodoList @update-count="updateCount" />
  </div>
</template>

<script setup>
import { ref } from "vue";
import axios from "axios";
import TodoList from "../components/TodoList.vue";

const count = ref({ total: 0, completed: 0 });

function updateCount(payload) {
  count.value = payload;
}
</script>
