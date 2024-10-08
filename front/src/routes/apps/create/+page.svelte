<script lang="ts">
    import { Button } from "$lib/components/ui/button/index.js";
    import * as Card from "$lib/components/ui/card/index.js";
    import { Input } from "$lib/components/ui/input/index.js";
    import { Label } from "$lib/components/ui/label/index.js";
    import ImageSelect from "$lib/ui/ImageSelect.svelte";
    import type { PageServerData } from "./$types";
    import type { Image, SearchedImage, AppParams } from "$lib/index.ts";
    import { goto } from "$app/navigation";

    export let data: PageServerData;
    let appName = "";
    let selectedImage = "";
    let defaultImages: Image[] = data.results;
    let searchedImages: SearchedImage[] = [];

    const handleImageSelect = (e: CustomEvent) => {
        selectedImage = e.detail;
    };
    const createApp = async (appParams: AppParams): Promise<any> => {
        const formData = new URLSearchParams();
        formData.append("appName", appParams.appName);
        formData.append("imageName", appParams.imageName);
        const req = new Request("?/createApp", {
            method: "POST",
            headers: {
                "Content-Type": "application/x-www-form-urlencoded",
            },
            body: formData.toString(),
        });
        const res = await fetch(req);
        if (!res.ok) throw new Error(`Error: ${res.statusText}`);
        return await res.json();
    };

    async function handleSubmit() {
        const res = await createApp({
            appName: appName,
            imageName: selectedImage,
        });
        console.log(res.status);
        if (res.status === 303) {
            goto("/apps");
        }
    }
</script>

<main class="flex container justify-center pt-12 h-full">
    <Card.Root class="w-[750px] border border-gray-800">
        <Card.Header>
            <Card.Title>Create project</Card.Title>
            <Card.Description
                >Deploy your new project in one-click.</Card.Description
            >
        </Card.Header>
        <Card.Content>
            <form on:submit={handleSubmit}>
                <div class="grid w-full items-center gap-4">
                    <div class="flex flex-col space-y-1.5">
                        <Label for="appName">Name</Label>
                        <Input
                            bind:value={appName}
                            id="appName"
                            name="appName"
                            type="text"
                            placeholder="Name of your project"
                            class="focus:outline-none"
                        />
                    </div>

                    <div class="flex flex-col space-y-1.5">
                        <Label for="imageSelect">Image</Label>
                        <ImageSelect
                            bind:value={selectedImage}
                            {defaultImages}
                            {searchedImages}
                            on:imageSelected={handleImageSelect}
                        />
                    </div>

                    <Card.Footer class="flex justify-between p-0">
                        <Button variant="outline">Cancel</Button>
                        <Button type="submit">Create</Button>
                    </Card.Footer>
                </div>
            </form>
        </Card.Content>
    </Card.Root>
</main>
