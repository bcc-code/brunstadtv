<script lang="ts" setup>
import { RouteLocationRaw } from 'vue-router';

defineProps<{routes: {
    title: string;
    route: RouteLocationRaw;
}[]}>()
</script>
<template>
    <div class="flex justify-center">
        <div class="max-w-lg">
            <router-link v-for="r in routes" :to="r.route">
                {{r.title}}
            </router-link>
        </div>
    </div>
</template>