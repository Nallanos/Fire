<script lang="ts">
  import { onMount } from "svelte";

  import { getApp, type App, deleteApp } from "$lib/api";
  import { page } from "$app/stores";
  let app: App;
  async function deletingApplication() {
    deleteApp($page.params?.id);
  }
  onMount(async () => {
    app = await getApp($page.params?.id);
  });
</script>

<main class="flex">
  <div class="bg-gray-950 rounded-lg flex flex-col border border-gray-800 p-6">
    <div class="flex items-center">
      <div class="text-xl font-bold">{app?.name}</div>
      <div class="bg-red-500 rounded-full w-[10px] h-[10px] ml-auto"></div>
    </div>
    <ul class="flex flex-col">
      <li class="flex gap-3">
        {app?.id}
      </li>
      <li class="text-gray-200">
        Listening on port
        <span>{app?.port}</span>
      </li>
    </ul>
    <button
      class="bg-red-600 rounded-md w-[80px] h-[30px] py-10px"
      on:click={deletingApplication}
    >
      Delete
    </button>
  </div>
</main>
