<script lang="ts">
    import "../app.css";
    import type { LayoutData } from "./$types";
    import { Toaster } from "svelte-sonner";

    export let data: LayoutData;
</script>

<Toaster position="top-center" />
<div class="grid grid-cols-[auto,1fr] gap-4 p-10 align-top">
    <div>
        <slot />
    </div>
    {#if data.favorites}
        <div class="grid-rows grid content-baseline gap-4">
            <h2 class="mb-4 text-3xl font-semibold text-gray-800">
                Favourites
            </h2>
            {#each [...data.favorites] as { mal_id, title, image } (mal_id)}
                <a
                    href="/{mal_id}"
                    class="block w-60 overflow-hidden rounded-lg bg-white no-underline shadow-lg transition-shadow duration-300 hover:shadow-xl"
                >
                    <img
                        src={image}
                        alt={title}
                        class="h-auto w-auto object-contain"
                    />
                    <div class="p-4">
                        <h4 class="text-xl font-semibold text-gray-800">
                            {title}
                        </h4>
                    </div>
                </a>
            {/each}
        </div>
    {/if}
</div>
