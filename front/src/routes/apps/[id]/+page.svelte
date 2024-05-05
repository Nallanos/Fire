<script lang="ts">
  import StatusIcon from "$lib/ui/StatusIcon.svelte";
  import {
    getActiveDeployment,
    getApp,
    type App,
    type Deployment,
  } from "$lib/api/apps";
  import { onMount } from "svelte";
  import { page } from "$app/stores";
  let id: string = $page.params?.id;
  let isActiveDeployment: boolean;
  let app: App;
  let status: string;
  let deployment: Deployment;
  let created_at: string;
  onMount(async () => {
    deployment = await getActiveDeployment(id);
    app = await getApp(id);
    if (deployment != undefined) {
      isActiveDeployment = true;
    }
    console.log(app, deployment);
    status = app.status;
    created_at = new Date(deployment.created_at).toLocaleDateString();
  });
</script>

{#if isActiveDeployment}
  <div
    class="bg-gray-950 rounded-lg flex flex-col border border-gray-800 p-6 container mx-auto"
  >
    <div class="flex items-center gap-2">
      <div class="text-xl font-bold">{deployment.id}</div>
      <div>
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
        Created the
        <p class="px-2 font-medium">
          {created_at}
        </p>
      </li>
    </ul>
  </div>
{/if}
