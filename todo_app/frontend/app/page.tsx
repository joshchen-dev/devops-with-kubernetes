import Image from "next/image";

export default function Home() {
  return (
    <div>
      <p className="text-3xl font-bold text-center my-5">Todo App</p>
      <Image src="http://localhost:8080/image" alt="test" width={600} height={600} className="mx-auto rounded-xl shadow-md shadow-mauve-800/80" />
      <p className="text-sm text-center my-3">devops with kubernetes 2026</p>
    </div>
  )
}
