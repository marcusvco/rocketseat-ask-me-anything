interface RemoveMessageReactionRequest {
    roomId: string
    messageId: string
}

export async function removeMessageReaction({ roomId, messageId }: RemoveMessageReactionRequest) {
    const response = await fetch(`${import.meta.env.VITE_APP_API_URL}/rooms/${roomId}/messages/${messageId}/react`, {
        method: "DELETE",
    })

    const data: { id: string } = await response.json()

    return { roomId: data.id }
}