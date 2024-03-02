<script lang="ts" setup>
import Update from '@/components/Update.vue';
import VersionDisplayer from '@/components/VersionDisplayer.vue';
import DockMenu from '@/components/menus/DockMenu.vue';
import GameMenu from '@/components/menus/GameMenu.vue';
import Title from '@/components/Title.vue';

</script>
<template>
    <Title />
    <div class="flex flex-column align-items-center">
        <GameMenu />
        <DockMenu />
        <VersionDisplayer />
        <Update />
    </div>
</template>