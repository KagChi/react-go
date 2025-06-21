import { usePage } from '@inertiajs/react';

export default function Home() {
    const { text } = usePage().props as unknown as { text: string };

    return (
        <div>
            Home: {text}
        </div>
    );
}
