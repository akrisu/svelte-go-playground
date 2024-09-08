<script lang="ts">
    import type { RecommendedAnime } from "../../routes/+page.server";

    export let subRecommendations: RecommendedAnime["entry"];
</script>

<div class="grid grid-flow-col gap-8">
    {#each subRecommendations as { title, mal_id, images }}
        <div>
            <h1>{title}</h1>

            <a href="/{mal_id}">
                <img
                    class="hover:cursor-pointer hover:opacity-50"
                    src={images.webp.large_image_url}
                    alt={title}
                    height="200px"
                    width="200px"
                />
            </a>
        </div>
    {/each}
</div>
