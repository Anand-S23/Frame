<script lang="ts">
	import { superForm } from "sveltekit-superforms/client";
	import type { PageData } from "./$types";
	import { z } from "zod";

    import { Input } from "$lib/components/ui/input";
    import { Button } from "$lib/components/ui/button";

	export let data: PageData;

    const loginSchema = z.object({
        email: z.string({ required_error: 'Email is required' }),
        password: z.string({ required_error: 'Password is required' }).trim()
    });

	const { form, errors, enhance, constraints } = superForm(data.form, {
		taintedMessage: "Are you sure you want leave?",
		validators: loginSchema
	});
</script>

<div class="flex flex-col w-full max-w-md gap-1.5 p-5">
	<form method="POST" use:enhance>
		<Input 
            type="email" id="email" name="email" placeholder="Email" bind:value={$form.email}
            class="mt-2 {$errors.email ? 'border-red-500' : ''}"
        />
		{#if $errors.email}
            <p class="text-sm text-red-500 mx-2">{$errors.email}</p>
		{/if}

		<Input 
            type="password" id="password" name="password" placeholder="Password" bind:value={$form.password}
            class="mt-2 {$errors.password ? 'border-red-500' : ''}"
        />
		{#if $errors.password}
            <p class="text-sm text-red-500 mx-2">{$errors.password}</p>
		{/if}

		<Button type="submit" class="mt-2">Submit</Button>
	</form>
</div>

