import { useSuspenseQuery } from "@tanstack/react-query";
import { useParams } from "react-router-dom";
import { useMessageWebsockets } from "../hooks/use-messages-websockets";
import { getRoomMessages } from "../http/get-room-messages";
import { Message } from "./message";

export function Messages() {
    const { roomId } = useParams()

    if (!roomId) {
        throw new Error("Componentes de mensagem devem ser usados na página de sala")
    }

    const { data } = useSuspenseQuery({
        queryKey: ["messages", roomId],
        queryFn: () => getRoomMessages({ roomId }),
    })

    useMessageWebsockets({ roomId })

    const sortedMessages = data.messages.sort((a, b) => {
        return b.amountOfReactions - a.amountOfReactions
    })

    return (
        <ol className="list-decimal list-outside px-3 space-y-8 ">
            {sortedMessages.map(message => {
                return (
                    <Message
                        key={message.id}
                        id={message.id}
                        text={message.text}
                        amountOfReactions={message.amountOfReactions}
                        answered={message.answered}
                    />
                )
            })}
        </ol>
    )
}