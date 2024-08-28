<script lang="ts">
  import { type Deployment } from "$lib/api/deployment/deploymentType";
  import StatusIcon from "../StatusIcon.svelte";

  export let deployment: Deployment;

  let status = deployment.status;
  let created_at = new Date(deployment.created_at).toLocaleDateString();
</script>

<div
  class="bg-[#09090b] rounded-lg flex flex-col border border-gray-800 p-6 container mx-auto"
>
  <div class="flex items-center gap-2">
    <div class="text-xl font-bold">{deployment.id}</div>
    <div>
      <StatusIcon {status} />
    </div>
  </div>

  <ul class="flex flex-col">
    <li class="flex gap-3"></li>
    <li class="flex">
      Created the
      <p class="px-2 font-medium">
        {created_at}
      </p>
    </li>
  </ul>
</div>
