import { getApp } from "$lib/api";
import type { LayoutLoad } from "./$types.js";

export const load: LayoutLoad = async (LayoutParams) => {
  return {
    app: await getApp(LayoutParams.params.id),
  };
};
