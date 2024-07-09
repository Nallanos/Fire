import { getApp } from "$lib/api/apps/getApp.ts";
import type { LayoutLoad } from "./$types.js";
export const load: LayoutLoad = async (LayoutLoad) => {
  return {
    app: await getApp(LayoutLoad.params.id),
  };
};
