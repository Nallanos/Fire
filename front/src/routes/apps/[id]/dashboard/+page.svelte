<script lang="ts">
    import Dashboard from "lib/ui/deployment/Dashboard.svelte";
    import type { PageData } from "./$types";

    export let data: PageData;
    const token = data.token;
    const app = data.app;
</script>

{#if token != undefined}
    <Dashboard {token} {app}></Dashboard>
{/if}
