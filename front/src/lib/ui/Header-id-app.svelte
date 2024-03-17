<script lang="ts">
  import { onMount } from "svelte";
  import { type App, deleteApp, getApp } from "$lib/api";
  import { page } from "$app/stores";
  let app: App;
  async function deletingApplication() {
    deleteApp($page.params?.id);
  }
  onMount(async () => {
    app = await getApp($page.params?.id);
  });
</script>

<header class="flex border-b border-gray-800 py-4 px-4">
  <div
    class="bg-gray-950 rounded-lg p-6 flex justify-between items-center w-full container mx-auto"
  >
    <h1 class="text-2xl font-bold">{app?.name}</h1>
    <ul class="flex gap-4">
      <button
        class="border rounded-md border-gray-800 w-[80px] h-[40px] py-10px"
        >Domain</button
      >

      <a href="/apps">
        <button
          class="bg-red-600 rounded-md w-[118px] h-[40px] py-10px"
          on:click={deletingApplication}
        >
          Delete
        </button>
      </a>
    </ul>
  </div>
</header>
