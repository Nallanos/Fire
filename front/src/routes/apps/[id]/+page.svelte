<script lang="ts">
  import type { Deployment } from "$lib/api/deployment/deploymentType";
  import StatusIcon from "$lib/ui/StatusIcon.svelte";
  import DeploymentCard from "$lib/ui/ApplicationPage/DeploymentCard.svelte";
  import type { LayoutData } from "./$types";
  import Button from "lib/components/ui/button/button.svelte";
  export let data: LayoutData;
  const app = data.app;
  let deploymentList: Deployment[] = data.deployments;
  let activeDeployment: Deployment = data.activeDeployment;
  let appError = data.app.error;
  let isActiveDeployment: boolean = !!activeDeployment;
  let status: string = activeDeployment?.status;
  let created_at: string = activeDeployment?.created_at;
</script>

<main class="grid gap-6 container mx-auto pt-4">
  {#if isActiveDeployment}
    <div
      class="bg-[#09090b] rounded-lg flex flex-col border border-gray-800 p-6 md:container md:mx-auto"
    >
      <div class="flex flex-col md:flex-row gap-2 justify-between">
        <div class="flex items-center gap-6">
          <p class="text-xl font-bold">{activeDeployment.id}</p>
          <StatusIcon {status} />
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
      <div class="gap-6 flex pt-4">
        <form method="POST" action={`/apps/${app.id}?/startContainer`}>
          <button class="w-20 h-8 bg-white text-black rounded-md">Start</button>
        </form>
        <form method="POST" action="?/stopContainer">
          <button class="w-20 h-8 border border-gray-800 rounded-md"
            >Stop</button
          >
        </form>
      </div>
    </div>
  {/if}

  <h1 class="font-bold text-lg">Deployment list</h1>
  <div class="grid gap-4 pb-4">
    {#if deploymentList != undefined}
      {#each deploymentList as deployment}
        <DeploymentCard {deployment}></DeploymentCard>
      {/each}
    {/if}
  </div>

  {#if appError}
    <p class="container mx-auto">Error loading app data: {appError}</p>
  {/if}
</main>
