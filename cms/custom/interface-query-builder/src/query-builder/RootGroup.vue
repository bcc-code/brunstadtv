<template>
	<div v-if="group">
        <h1>Filter</h1>
		<Group 
            style="margin-left: 10px"
			:value="group"
			:fields="fields" 
			@delete="group = null"
            @update:value="(f) => group = f"
            @change="handleChange"
		/>
        <hr class="separator"/>
        <h1>Sort by</h1>
		<div class="sort-by">
            <select v-model="sortBy" @change="handleChange">
                <option
                    v-for="field in fields"
                    :value="field.name"
                >{{snakeToPascal(field.name)}}
                </option>
            </select>
            <h3>Direction</h3>
			<select v-model="sortByDirection" @change="handleChange">
                <option value="asc">Ascending</option>
                <option value="desc">Descending</option>
			</select>
		</div>
	</div>
	<div v-else>
		<VButton @click="clearGroup()">Create filter</VButton>
	</div>
</template>

<script lang="ts" setup>
import { v4 as uuid } from "uuid";
import { snakeToPascal, Field as TField, Root } from ".";
import Group from "./Group.vue";
import { ref } from "vue";
import VButton from "./VButton.vue";

const props = defineProps<{
    value: Root | null;
    fieldFactory: () => Promise<TField[]>
}>()
const emit = defineEmits<{
	(e: "update:value", value: Root | null),
}>()

const fields = ref(await props.fieldFactory() as TField[]);
const group = ref(props.value?.filter ?? null);
const sortBy = ref(props.value?.sortBy ?? null);
const sortByDirection = ref(props.value?.sortByDirection ?? null)

function handleChange(): void {
    emit('update:value', {
        id: uuid(),
        filter: group.value,
        sortBy: sortBy.value,
        sortByDirection: sortByDirection.value,
    })
}

function clearGroup() {
	group.value = {
        "and": [],
    };
    handleChange();
}
</script>
<style>
h1 {
    font-size: 18px;
    font-weight: 700;
}

.sort-by {
    padding: 10px;
}

.separator {
    margin: 10px;
    border-color: #31363d;
}
</style>
