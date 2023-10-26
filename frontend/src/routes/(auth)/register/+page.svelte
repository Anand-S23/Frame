<script lang="ts">
	import { superForm } from "sveltekit-superforms/client"
	import type { PageData } from "./$types"
	import { z } from "zod"

    import { Input } from "$lib/components/ui/input";
    import { Button } from "$lib/components/ui/button";

	export let data: PageData;

    const registerSchema = z.object({
        username: z
            .string({ required_error: 'Username is required' })
            .min(5, { message: "Username must be at least 5 characters" })
            .max(25, { message: "Username must be less than 25" })
            .trim(),
        email: z
            .string({ required_error: 'Email is required' })
            .max(64, { message: 'Email must be less than 64 characters' })
            .email({ message: 'Email must be a valid email address' }),
        password: z
            .string({ required_error: 'Password is required' })
            .min(8, { message: 'Password must be at least 8 characters' })
            .max(32, { message: 'Password must be less than 32 characters' })
            .trim(),
    });

	const { form, errors, enhance, constraints } = superForm(data.form, {
		taintedMessage: "Are you sure you want leave?",
		validators: registerSchema
	})
</script>

<div class="flex flex-col w-full max-w-md gap-1.5 p-5">
	<form method="POST" use:enhance>
		<Input 
            type="text" id="username" name="username" placeholder="Username" bind:value={$form.username}
            class="mt-2 {$errors.username ? 'border-red-500' : ''}"
        />
		{#if $errors.username}
            <p class="text-sm text-red-500 mx-2">{$errors.username}</p>
		{/if}

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
