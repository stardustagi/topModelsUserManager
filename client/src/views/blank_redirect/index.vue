<template>

</template>

<script setup lang="ts">
import { syncUserStoreFromStorage } from '@/utils/auth';
import { onMounted } from 'vue';
import { useRouter } from "vue-router";

const router = useRouter();

onMounted(async () => {
    syncUserStoreFromStorage();
    router.replace("/console/person_center/keymanager");
})
</script>



<style lang="scss" scoped></style>