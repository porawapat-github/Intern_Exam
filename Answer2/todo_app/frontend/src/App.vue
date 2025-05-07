<template>
  <div class="min-h-screen bg-gray-100 text-gray-800">
    <nav class="bg-white shadow-md">
      <div class="max-w-4xl mx-auto px-4 py-3 flex justify-between items-center">
        <h1 class="text-xl font-bold text-blue-600">📝 Todo App</h1>
        <div class="space-x-4 text-sm">
          <router-link to="/" class="text-gray-700 hover:text-blue-600 font-medium">
            หน้าหลัก
          </router-link>
          <router-link to="/history" class="text-gray-700 hover:text-blue-600 font-medium">
            ประวัติงานที่เสร็จ
          </router-link>
        </div>
      </div>
    </nav>

    <main class="max-w-4xl mx-auto py-6 px-4">
      <router-view />
    </main>
  </div>
</template>

<script setup>
// ไม่มี logic พิเศษ
</script>

<style>
/* Optional: เปลี่ยน font ทั้งระบบ */
body {
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}
</style>
