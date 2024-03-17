<script lang="ts">
  import { onMount } from "svelte";
  import { getApp, type App } from "$lib/api";
  import { page } from "$app/stores";
  import HeaderIdApp from "$lib/ui/Header-id-app.svelte";
  import QuickSetup from "$lib/ui/Quick-Setup.svelte";
  let app: App;
  onMount(async () => {
    app = await getApp($page.params?.id);
  });
</script>

<main>
  <HeaderIdApp></HeaderIdApp>
  <QuickSetup></QuickSetup>
</main>
