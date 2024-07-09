import { listApps } from "$lib/api/apps/listApps";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ params }) => {
  return {
    apps: await listApps(),
  };
};
