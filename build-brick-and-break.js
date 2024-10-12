export function build(brickCount){
    let counter= 1
    
    const intervalid = setInterval(()=>{
        if (counter>brickCount){
            clearInterval(intervalid)
            return
        }
        const newBrick=document.createElement('div');
        newBrick.id=`brick-${counter}`

        if (counter===Math.ceil(brickCount/2)){
            newBrick.setAttribute(`data-foundation`,`true`)
        }

        document.body.appendChild(newBrick)

        counter++
    },100)

}

export function repair(...ids){
    ids.forEach(id=>{
        const brick = document.getElementById(id)
     
    if (brick){
        const isMiddleColummn = brick.getAttribute(`data-foundation`)===`true`

        brick.setAttribute(`data-repaired`,isMiddleColummn ? `in progress` :`true`)
    }})}

    export function destroy(){
        const bricks = document.querySelectorAll(`div[id^="brick-"]`)
        const lastBrick = bricks[bricks.length-1]

        if (lastBrick){
            lastBrick.remove()
        }
    }
