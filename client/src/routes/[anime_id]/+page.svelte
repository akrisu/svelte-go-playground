<script lang="ts">
    import { enhance } from "$app/forms";
    import { toast } from "svelte-sonner";
    import type { PageData } from "./$types";

    export let data: PageData;
</script>

<div class="grid grid-flow-col gap-8">
    <div class="relative w-full max-w-sm overflow-hidden rounded-lg shadow-lg">
        <img
            src={data.anime.images.webp.large_image_url}
            alt={data.anime.title}
            class="h-auto w-full"
        />
        <div
            class="absolute bottom-0 left-0 w-full bg-gradient-to-t from-black to-transparent p-4 text-white"
        >
            <h2 class="text-lg font-semibold shadow-lg">{data.anime.title}</h2>
        </div>
    </div>
    <div>
        <p class="text-base leading-relaxed text-gray-700">
            {data.anime.synopsis}
        </p>
        <form
            action="?/addToFavorites"
            method="post"
            use:enhance={() =>
                async ({ result, update }) => {
                    console.log(result);
                    if (result.type === "success" && result.data?.["success"]) {
                        toast.success("Favourite has been added successfully");
                    }
                    if (
                        result.type === "success" &&
                        !result.data?.["success"]
                    ) {
                        toast.error(`${result.data?.["message"]}`);
                    }
                    update();
                }}
        >
            <input type="hidden" name="mal_id" value={data.anime.mal_id} />
            <input type="hidden" name="title" value={data.anime.title} />
            <input
                type="hidden"
                name="image"
                value={data.anime.images.webp.image_url}
            />
            <button class="mt-4 rounded bg-gray-300 p-4" type="submit">
                Add to favorites
            </button>
        </form>

        <form
            action="?/removeFromFavorites"
            method="post"
            use:enhance={() =>
                async ({ result, update }) => {
                    if (result.type === "success" && result.data?.["success"]) {
                        toast.success(
                            "Favourite has been removed successfully",
                        );
                    }
                    if (
                        result.type === "success" &&
                        !result.data?.["success"]
                    ) {
                        toast.error(`${result.data?.["message"]}`);
                    }
                    update();
                }}
        >
            <input type="hidden" name="mal_id" value={data.anime.mal_id} />
            <button class="mt-4 rounded bg-gray-300 p-4" type="submit">
                Remove from favorites
            </button>
        </form>

        <div class="mt-10">
            <a class="rounded bg-gray-300 p-4" href="/">Go back to list</a>
        </div>
    </div>
</div>
