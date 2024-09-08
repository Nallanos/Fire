<script lang="ts">
    import { page } from "$app/stores";
</script>

{#if $page.error != null}
    <h1>{$page.status}: {$page.error.message} caca</h1>
{/if}
