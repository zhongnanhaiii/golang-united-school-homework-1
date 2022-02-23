package solution

import (
    "github.com/kyokomi/emoji/v2"
)

func GetMesage () string {
    emojiSolution := emoji.Sprint("Hello :world_map:")
    return emojiSolution
}

