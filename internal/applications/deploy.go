package applications
	// slog.Info("Pulling image", "image", app.Image)
	// // Lis l'image docker
	// // reader, err := a.docker.ImagePull(context.Background(), app.Image, types.ImagePullOptions{})
	// // if err != nil {
	// // 	slog.Error("Error pulling image", "err", err)
	// // 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// // 	return
	// // }

	// // slog.Info("Copying image pull output")
	// // if _, err := io.Copy(os.Stdout, reader); err != nil {
	// // 	slog.Error("Error copying image pull output", "err", err)
	// // 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// // 	return
	// // }

	// slog.Info("Creating container", "image", app.Image)
	// // On crée un container, en initialisant l'image du container avec app.Image
	// // Puis, on lance le container
	// container, err := a.docker.ContainerCreate(context.Background(), &container.Config{
	// 	Image: app.Image,
	// }, &container.HostConfig{
	// 	PortBindings: nat.PortMap{
	// 		nat.Port("8081/tcp"): []nat.PortBinding{
	// 			{
	// 				HostPort: "3000",
	// 				HostIP:   "0.0.0.0",
	// 			},
	// 		},
	// 	},
	// }, nil, nil, app.ID)

	// if err != nil {
	// 	slog.Error("Error creating container", "err", err)
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// slog.Info("Starting container", "id", container.ID)
	// err = a.docker.ContainerStart(context.Background(), container.ID, types.ContainerStartOptions{})
	// if err != nil {
	// 	slog.Error("Error starting container", "err", err)
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// slog.Info("Container started", "id", container.ID)

	// w.Header().Set("Content-Type", "application/json")

