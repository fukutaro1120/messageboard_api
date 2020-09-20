package controller

import "github.com/gin-gonic/gin"

type PingController struct {
	result string // pong とはいってる
}

func (p *PingController) Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

// class PingController {
// 	private $result = "pong";
// 	public function Ping(ctx) {
// 		ctx->JSON(200, gin.H){
// 			"message": $this->result
// 		}
// 	} 
// }
