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

  let style = "bg-white";
  let buttonText = "";
  let visible = false;

  async function deletingApplication() {
    deleteApp($page.params?.id);
  }

  async function deployingApplication() {
    await deployApp($page.params?.id).then(
      () => {
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
        style = "bg-red-400";
      }
    );
  }

  async function refreshApp() {
    app = await getApp($page.params?.id);
  }
  function refreshButtonText() {
    refreshApp().finally(() => {
      if (app.status == "inactive") {
        buttonText = "Deploy";
      } else if (app.status == "completed" || "pending") {
        buttonText = "Stop";
      }
    });
  }

  async function DeployAndStopButtonHandler() {
    if (app.status == "inactive") {
      await deployingApplication().finally(() => {
        refreshButtonText();
      });
    }
    if (app.status == "completed" || app.status == "error") {
      await stopContainer($page.params?.id).finally(() => {
        refreshButtonText();
      });
    }
  }
  onMount(async () => {
    app = await getApp($page.params?.id);
    refreshButtonText();
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
        >{buttonText}</button
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
