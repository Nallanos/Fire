<script lang="ts">
  import { onMount } from "svelte";
  import {
    type App,
    deleteApp,
    getApp,
    deployApp,
    stopContainer,
  } from "$lib/api";
  import { page } from "$app/stores";
  import { fade } from "svelte/transition";
  let app: App;

  async function deletingApplication() {
    deleteApp($page.params?.id);
  }

  let status = "Deploy";
  let style = "bg-white";
  let visible = false;

  async function deployingApplication() {
    await deployApp($page.params?.id).then(
      () => {
        status = "Stop";
        style = "bg-white";
        visible = true;
        setTimeout(() => {
          visible = false;
        }, 2000);
      },
      () => {
        visible = true;
        setTimeout(() => {
          visible = false;
        }, 2000);
        status = "Failed";
        style = "bg-red-400";
      }
    );
  }

  async function DeployAndStopButtonHandler() {
    if (status == "Deploy") {
      deployingApplication();
    } else if (status == "Stop") {
      await stopContainer($page.params?.id).finally(() => {
        status = "Deploy";
      });
    }
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
      {#if visible}
        <p
          transition:fade={{ duration: 2000 }}
          class="w-[170px] bg-green-400 rounded-md items-center justify-center flex text-black"
        >
          Application deployed !
        </p>
      {/if}
      <button
        on:click={DeployAndStopButtonHandler}
        class={`${style} border rounded-md border-gray-800 w-[80px] h-[40px] py-10px text-black`}
        >{status}</button
      >
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
