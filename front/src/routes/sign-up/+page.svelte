<script lang="ts">
    import { Button } from "$lib/components/ui/button/index.js";
    import * as Card from "$lib/components/ui/card/index.js";
    import { Input } from "$lib/components/ui/input/index.js";
    import { Label } from "$lib/components/ui/label/index.js";
    let email = "";
    let password = "";
    let username = "";
</script>

<main class="flex container flex-col justify-center items-center pt-12 h-full">
    <h1 class="text-3xl font-bold pb-6">Fire</h1>
    <Card.Root class="border md:w-[500px] border border-gray-800  ">
        <Card.Header>
            <Card.Title>Sign Up</Card.Title>
        </Card.Header>
        <Card.Content>
            <form method="POST">
                <div class="grid w-full items-center gap-4 border-none">
                    <div class="flex flex-col space-y-1.5">
                        <div class="pb-8">
                            <Label for="name">Username</Label>
                            <Input
                                bind:value={username}
                                name="username"
                                id="username"
                                class="focus:outline-none"
                                required
                            />
                        </div>
                        <Label for="name">Email</Label>
                        <Input
                            bind:value={email}
                            name="email"
                            type="email"
                            id="email"
                            class="focus:outline-none"
                            required
                        />
                        <Label for="name">Password</Label>
                        <Input
                            bind:value={password}
                            name="password"
                            id="password"
                            type="password"
                            class="focus:outline-none"
                            required
                        />
                    </div>

                    <Card.Footer class="flex justify-between p-0">
                        <Button type="submit">Sign Up</Button>
                    </Card.Footer>
                </div>
            </form>
        </Card.Content>
    </Card.Root>
</main>
