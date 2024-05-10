<script lang="ts">
  import StatusIcon from "$lib/ui/StatusIcon.svelte";
  import {
    StartContainer,
    getActiveDeployment,
    getApp,
    listDeployments,
    stopContainer,
    type App,
    type Deployment,
  } from "$lib/api/apps";
  import { onMount } from "svelte";
  import { page } from "$app/stores";
  import DeploymentCard from "$lib/ui/ApplicationPage/DeploymentCard.svelte";
  let id: string = $page.params?.id;
  let isActiveDeployment: boolean;
  let app: App;
  let status: string;
  let deployment: Deployment;
  let created_at: string;
  let deploymentList: Deployment[] = [];

  async function StartActiveDeployment() {
    await StartContainer(id);
  }
  async function StopActiveDeployment() {
    await stopContainer(id);
  }

  onMount(async () => {
    deploymentList = await listDeployments(id);
    deployment = await getActiveDeployment(id);
    app = await getApp(id);
    status = app.status;
    created_at = new Date(deployment.created_at).toLocaleDateString();
    if (deployment.id == undefined) {
      isActiveDeployment = false;
    } else {
      isActiveDeployment = true;
    }
    console.log(deploymentList);
  });
</script>

<main class="grid gap-6">
  {#if isActiveDeployment == true}
    <div
      class="bg-gray-950 rounded-lg flex flex-col border border-gray-800 p-6 container mx-auto"
    >
      <div class="flex items-center gap-2 justify-between">
        <div class="flex items-center gap-6">
          <p class="text-xl font-bold">{deployment.id}</p>
          <StatusIcon {status} />
        </div>
        <div class="gap-4 flex">
          <button
            class="w-20 h-8 bg-white text-black rounded-md"
            on:click={StartActiveDeployment}>Start</button
          >
          <button
            class="w-20 h-8 border border-gray-800 rounded-md"
            on:click={StopActiveDeployment}>Stop</button
          >
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
      </ul>
    </div>
  {/if}
  <h1 class="container mx-auto font-bold text-lg">Deployment list</h1>
  <div class="grid gap-4 container mx-auto">
    {#each deploymentList as deployment}
      <DeploymentCard {deployment}></DeploymentCard>
    {/each}
  </div>
</main>
