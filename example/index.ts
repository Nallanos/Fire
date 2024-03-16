Bun.serve({
  fetch(req: Request): Response | Promise<Response> {
    return new Response("Hello World!");
  },
  port: 8081,
});
