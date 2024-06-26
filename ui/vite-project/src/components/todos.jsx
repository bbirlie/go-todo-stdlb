import { useQuery } from "@tanstack/react-query"

function Todos() {
    const { data, isLoading } = useQuery({
        queryKey: ["todos"],
        queryFn: async () => {
            try {
                const res = await fetch("http://localhost:8080/todos")
                const data = await res.json()

                if (!res.ok) {
                    throw new Error(data.error || "rip")
                }
                return data || []
            } catch (error) {
                console.log(error);
            }
        }
    })
    return (
        <>
            <h1>Todos</h1>
            {isLoading && <p>loading</p>}
            {!isLoading && <ul>
                {data.map(todo => (
                    <li key={todo.id}>
                        {todo.title} + {todo.text}
                    </li>
                ))}
            </ul>}
        </>
    )
}
export default Todos