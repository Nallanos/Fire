<script lang="ts">
  import type { App } from "$lib/api/apps/appTypes";
  import type { Deployment } from "$lib/api/deployment/deploymentType";
  import StatusIcon from "$lib/ui/StatusIcon.svelte";
  import { listDeployments } from "$lib/api/deployment/listDeployments";
  import { getActiveDeployment } from "$lib/api/deployment/getActiveDeployment";
  import { onMount } from "svelte";
  import { page } from "$app/stores";
  import DeploymentCard from "$lib/ui/ApplicationPage/DeploymentCard.svelte";
  import { getApp } from "$lib/api/apps/getApp";
  let id: string = $page.params?.id;
  let isActiveDeployment: boolean = false;
  let app: App;
  let status: string;
  let activeDeployment: Deployment;
  let created_at: string;
  let deploymentList: Deployment[] = [];
  let deploymentListError: string | null = null;
  let activeDeploymentError: string | null = null;
  let appError: string | null = null;

  onMount(async () => {
    try {
      const appResult = await getApp(id);
      if (appResult.error) {
        appError = appResult.error;
      } else {
        app = appResult.data;
        status = app.status;
        console.log(app);
      }
      const deploymentResult = await listDeployments(id);
      if (deploymentResult.error) {
        deploymentListError = deploymentResult.error;
        deploymentList = [];
      } else {
        deploymentList = deploymentResult.data;
      }

      const activeDeploymentResult = await getActiveDeployment(id);
      if (activeDeploymentResult.error) {
        activeDeploymentError = activeDeploymentResult.error;
        isActiveDeployment = false;
      } else {
        activeDeployment = activeDeploymentResult.data;
        created_at = new Date(activeDeployment.created_at).toLocaleDateString();
        isActiveDeployment = true;
      }
    } catch (err) {
      console.error("Unexpected error:", err);
    }
  });
</script>

<main class="grid gap-6">
  {#if isActiveDeployment && !activeDeploymentError}
    <div
      class="bg-gray-950 rounded-lg flex flex-col border border-gray-800 p-6 container mx-auto"
    >
      <div class="flex items-center gap-2 justify-between">
        <div class="flex items-center gap-6">
          <p class="text-xl font-bold">{activeDeployment.id}</p>
          <StatusIcon {status} />
        </div>
        <div class="gap-4 flex">
          <form method="POST" action="?/startContainer">
            <button class="w-20 h-8 bg-white text-black rounded-md"
              >Start</button
            >
          </form>
          <form method="POST" action="?/stopContainer">
            <button class="w-20 h-8 border border-gray-800 rounded-md"
              >Stop</button
            >
          </form>
        </div>
      </div>

      <ul class="flex flex-col">
        <li class="flex gap-3"></li>
        <li class="text-gray-200">
          Listening on port
          <span class="font-medium px-2">{app.port}</span>
        </li>
        <li class="flex">
          Created the
          <p class="px-2 font-medium">
            {created_at}
          </p>
        </li>
        <li>
          <a
            href={`http://localhost:${app.port}`}
            class="text-blue-600 underline">{`localhost:${app.port}`}</a
          >
        </li>
      </ul>
    </div>
  {/if}

  <h1 class="container mx-auto font-bold text-lg">Deployment list</h1>
  <div class="grid gap-4 container mx-auto">
    {#if deploymentListError || deploymentList == undefined}
      <p>Error loading deployment list: {deploymentListError}</p>
    {:else}
      {#each deploymentList as deployment}
        <DeploymentCard {deployment}></DeploymentCard>
      {/each}
    {/if}
  </div>

  {#if activeDeploymentError}
    <p class="container mx-auto">
      Error loading active deployment: {activeDeploymentError}
    </p>
  {/if}

  {#if appError}
    <p class="container mx-auto">Error loading app data: {appError}</p>
  {/if}
</main>
