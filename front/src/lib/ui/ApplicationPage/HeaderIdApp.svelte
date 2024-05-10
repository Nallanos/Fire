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
  let buttonText: string;
  let visible = false;

  async function deletingApplication() {
    await deleteApp($page.params?.id);
    window.location.href = "/apps";
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
  async function DeployAndStopButtonHandler(): Promise<void> {
    if (app.status == "inactive") {
      await deployingApplication().finally(() => {
        refreshButtonText();
      });
    } else {
      await stopContainer($page.params?.id).finally(() => {
        refreshButtonText();
      });
    }
  }

  async function refreshButtonText(): Promise<void> {
    await refreshApp().finally(() => {
      console.log(app);
      if (app.status == "inactive") {
        buttonText = "Deploy";
        console.log("the status of app is inactive");
      } else if (app.status == "completed" || app.status == "pending") {
        buttonText = "Stop";
        console.log("the status of app is completed", buttonText);
      }
    });
  }
  onMount(async () => {
    app = await getApp($page.params?.id);
  });
  refreshButtonText();
</script>

<header class="flex">
  <div
    class="bg-gray-950 rounded-lg py-6 flex justify-between items-center w-full container mx-auto"
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
        class="bg-red-600 rounded-md w-[118px] h-[40px] py-10px"
        on:click={deletingApplication}
      >
        Delete
      </button>
    </ul>
  </div>
</header>
