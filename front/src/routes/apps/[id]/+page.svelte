<script lang="ts">
  import { onMount } from "svelte";
  import { getApp, type App } from "$lib/api";
  import { page } from "$app/stores";
  import HeaderIdApp from "$lib/ui/HeaderIdApp.svelte";
  import QuickSetup from "$lib/ui/QuickSetup.svelte";
  let app: App;
  onMount(async () => {
    app = await getApp($page.params?.id);
  });
</script>

<main>
  <HeaderIdApp></HeaderIdApp>
  <QuickSetup></QuickSetup>
</main>
