<script lang="ts">
    import {
        createSvelteTable,
        flexRender,
        getCoreRowModel,
        getPaginationRowModel,
        type ColumnDef,
    } from "@tanstack/svelte-table";

    import type { PageData } from "./$types";
    import type { RecommendedAnime } from "./+page.server";

    import { Animes } from "$lib/components";
    import { onMount } from "svelte";
    import { gsap, ScrollTrigger } from "$lib/gsap";

    export let data: PageData;
    let loading = true;

    const pageSize = 10;
    const columns: ColumnDef<RecommendedAnime>[] = [
        {
            accessorKey: "content",
            header: "Description",
        },
        {
            accessorKey: "entry",
            header: "Recommendations",
            cell: ({ row }) =>
                flexRender(Animes, { subRecommendations: row.original.entry }),
        },
    ];
    const table = createSvelteTable({
        columns,
        data: data.recommended.data,
        getPaginationRowModel: getPaginationRowModel(),
        getCoreRowModel: getCoreRowModel(),
    });

    $table.setPageSize(pageSize);

    const handlePageChange = (direction: "next" | "previous") => {
        direction === "next" && $table.nextPage();
        direction === "previous" && $table.previousPage();
        attachAnimation();
        document.body.scrollIntoView();
    };

    const attachAnimation = () => {
        ScrollTrigger.refresh();
        Array(pageSize)
            .fill(0)
            .map((_, i) => {
                gsap.from(`.row-${i}`, {
                    x: -200,
                    opacity: 0,
                    duration: 3,
                    ease: "power4.out",
                    scrollTrigger: {
                        trigger: `.row-${i}`,
                        start: "top 80%",
                    },
                });
            });
    };
    onMount(() => {
        attachAnimation();
        loading = false;
    });
</script>

{#if loading}
    <div
        class="absolute left-0 top-0 flex h-full w-full items-center justify-center bg-gray-900"
    >
        <svg
            class="mr-3 h-10 w-10 animate-spin text-white"
            xmlns="http://www.w3.org/2000/svg"
            fill="white"
            viewBox="0 0 24 24"
            stroke="currentColor"
        >
            <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M12 6v6m0 0v6m0-6h6m-6 0H6"
            />
        </svg>
    </div>
{/if}

<h2 class="mb-4 text-3xl font-semibold text-gray-800">Recommendations</h2>
<table class="w-full table-auto text-left">
    <thead class="border-b border-gray-300">
        {#each $table.getHeaderGroups() as headerGroup}
            <tr>
                {#each headerGroup.headers as header}
                    <th class="p-4">
                        {#if !header.isPlaceholder}
                            <svelte:component
                                this={flexRender(
                                    header.column.columnDef.header,
                                    header.getContext(),
                                )}
                            />
                        {/if}
                    </th>
                {/each}
            </tr>
        {/each}
    </thead>
    <tbody>
        {#each $table.getRowModel().rows as row, i}
            <tr class={`${i % 2 === 0 ? "bg-gray-100" : ""} row-${i}`}>
                {#each row.getVisibleCells() as cell, i}
                    <td class={`p-4 align-top ${i === 0 ? "max-w-md" : ""}`}>
                        <svelte:component
                            this={flexRender(
                                cell.column.columnDef.cell,
                                cell.getContext(),
                            )}
                        />
                    </td>
                {/each}
            </tr>
        {/each}
    </tbody>
</table>
<button
    class="mt-4 rounded bg-gray-300 p-4 hover:bg-gray-100 disabled:bg-gray-50"
    on:click={() => handlePageChange("previous")}
    disabled={!$table.getCanPreviousPage()}
>
    Prev
</button>

<button
    class="mt-4 rounded bg-gray-300 p-4 hover:bg-gray-200 disabled:text-gray-200"
    on:click={() => handlePageChange("next")}
    disabled={!$table.getCanNextPage()}
>
    Next
</button>

Page {$table.getState().pagination.pageIndex + 1} of {$table.getPageCount()}
