export type App = {
  id: string;
  name: string;
  image: string;
  port: number;
  status: string;
};

export type CreateAppParams = {
  name: string;
};
