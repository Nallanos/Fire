<script lang="ts">
    import { Button } from "$lib/components/ui/button/index.js";
    import * as Card from "$lib/components/ui/card/index.js";
    import * as Select from "$lib/components/ui/select/index.js";
    import { Input } from "$lib/components/ui/input/index.js";
    import { Label } from "$lib/components/ui/label/index.js";
    import { createApp } from "$lib/api/apps/createApp";
    let appName = "";
    let linkAppImage = "";
    async function updateInputVal(event: Event) {
        let inputElement = event.target as HTMLInputElement;
        appName = inputElement.value;
    }
</script>

<main class="flex container justify-center pt-12 h-full">
    <Card.Root class="w-[750px] border border-gray-800 ">
        <Card.Header>
            <Card.Title>Create project</Card.Title>
            <Card.Description
                >Deploy your new project in one-click.</Card.Description
            >
        </Card.Header>
        <Card.Content>
            <form method="post" action="?/createApp">
                <div class="grid w-full items-center gap-4">
                    <div class="flex flex-col space-y-1.5">
                        <Label for="name">Name</Label>
                        <Input
                            bind:value={appName}
                            id="appName"
                            name="appName"
                            type="appName"
                            placeholder="Name of your project"
                            class="focus:outline-none"
                        />
                    </div>
                    <div class="flex flex-col space-y-1.5">
                        <Label for="framework">Image</Label>
                        <Input
                            bind:value={linkAppImage}
                            id="linkAppImage"
                            name="linkAppImage"
                            type="linkAppImage"
                            placeholder="docker.io/library/nginx:latest"
                            class="focus:outline-none"
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
