<script>
    let plans = [
        {
            title: "Pay-as-you-go",
            price: "$5",
            description: "Base price (connect 2 servers)",
            details: [
                "Connect 2 servers (included in the base price)",
                "$3 per additional server (billed monthly)",
                "Deploy unlimited applications/databases/services",
                "Free email notifications",
                "Limited support by email",
            ],
            button: "Get Started in the Cloud",
        },
    ];
</script>

<div class="flex flex-col items-center min-h-screen text-white py-8">
    <h1 class="text-4xl font-bold mb-4">Pricing</h1>
    <p class="text-lg text-gray-300 mb-6">
        Pricing that adjusts according to your needs.
    </p>

    {#each plans as plan}
        <div class="bg-[#09090b] rounded-lg p-6 w-full max-w-lg mb-6 shadow-lg">
            <h2 class="text-2xl font-semibold text-red-500 mb-2">
                {plan.title}
            </h2>
            <p class="text-3xl font-bold mb-4">{plan.price}</p>
            <p class="text-gray-400 mb-4">{plan.description}</p>
            <ul class="list-none space-y-2 mb-6">
                {#each plan.details as detail}
                    <li class="flex items-center space-x-2">
                        <span class="text-green-400">✔</span>
                        <span>{detail}</span>
                    </li>
                {/each}
            </ul>
            <button
                class="w-full py-3 bg-red-600 hover:bg-red-700 text-white font-semibold rounded-md"
            >
                {plan.button}
            </button>
        </div>
    {/each}
</div>
