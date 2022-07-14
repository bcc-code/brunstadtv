<template>
    <div class="max-w-sm rounded overflow-hidden shadow-lg">
        <slot name="header"></slot>
        <div class="px-6 py-4">
            <slot></slot>
        </div>
        <div class="px-6 pt-4 pb-2">
            <slot name="footer"></slot>   
        </div>
    </div>
</template>
