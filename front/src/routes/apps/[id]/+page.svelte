<script lang="ts">
  import type { App } from "$lib/api/apps/appTypes";
  import type { Deployment } from "$lib/api/deployment/deploymentType";
  import StatusIcon from "$lib/ui/StatusIcon.svelte";
  import { page } from "$app/stores";
  import DeploymentCard from "$lib/ui/ApplicationPage/DeploymentCard.svelte";
  import type { PageData } from "./$types";
  export let data: PageData;

  let app: App = data.app.data;
  let deploymentList: Deployment[] = data.deploymentResult.data;
  let activeDeployment: Deployment = data.activeDeploymentResult.data;
  let deploymentListError = data.deploymentResult.error;
  let activeDeploymentError = data.activeDeploymentResult.error;
  let appError = data.app.error;
  let id: string = $page.params?.id;
  let isActiveDeployment: boolean = !!activeDeployment;
  let status: string = activeDeployment?.status;
  let created_at: string = activeDeployment?.created_at;
</script>

<main class="grid gap-6 container mx-auto">
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
          Created at
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

  <h1 class="font-bold text-lg">Deployment list</h1>
  <div class="grid gap-4">
    {#if deploymentListError}
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
