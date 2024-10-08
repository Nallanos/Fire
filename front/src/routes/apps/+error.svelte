<script lang="ts">
    import { page } from "$app/stores";
</script>

<main class="container flex mx-auto">
    <header class="flex flex-col gap-8">
        {#if $page.error != null}
            <h1>{$page.status}: {$page.error.message}</h1>

            {#if $page.status === 401}
                <h1>You're not logged, please login</h1>
            {/if}
        {/if}
    </header>
</main>
