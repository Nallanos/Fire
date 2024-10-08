<script lang="ts">
    import Logs from "lib/ui/deployment/Logs.svelte";
    import type { PageData } from "./$types";

    export let data: PageData;
    const token = data.token;
    const app = data.app;
</script>

{#if token && app != undefined}
    <Logs {token} {app} />
{/if}
