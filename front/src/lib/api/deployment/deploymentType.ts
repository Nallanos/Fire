export type finished_at = {
  Time: string;
  Valid: boolean;
};
export type Deployment = {
  id: string;
  app_id: string;
  status: string;
  created_at: string;
  finished_at: finished_at;
};
