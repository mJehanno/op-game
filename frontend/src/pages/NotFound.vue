<script setup lang="ts">
import { useRouter } from 'vue-router';

import Message from 'primevue/message';
import Button from 'primevue/button';


const router = useRouter();
</script>
<template>
    <div>
        <Message severity="error">Sorry, but this page is not implemented yet !</Message>
        <Button label="Back" class="m-4" @click="router.push('/')"></Button>
    </div>
</template>