package techtask
 
// Вводные:
 
type Data struct {
    ID      int
    Payload map[string]interface{}
}
 
type Reader interface {
    Read() []*Data
}
 
type Processor interface {
    Process(Data) ([]*Data, error)
}
 
type Writer interface {
    Write([]*Data)
}
 
// Необходимо имплементировать интерфейс Manager так, чтобы он:
// - перманентно и синхронно принимал данные из Reader
// - обрабатывал все полученные данные на каждом из списка Processor
// - в случае отсутствия ошибок при обработке, полученные в процессе обработки данные передавал в Writer
// - при возникновении ошибки при обработке, ничего не отправлял во Writer
 
type Manager interface {
    Manage() // blocking
}
 
// Реализация:

type ManageObject struct{
    r *Reader
    p []*Processor
    w * Writer
    ctx *context.Context
}

func (mo *ManageObject) Manage(){
    resDataSlc:=make([]*Data,len(dataSlc)*len(mo.p))
    
    var err  error
    
    for{
        
    dataSlc:=mo.r.Read()
    wg:=&sync.WaitGroup{}
    ctxC,myCancel:= context.WithCancel()
    chProc:= make(chan *Data,len(mo.p))
    chRes:= make(chan *Data,len(mo.p))
           
    for k,_:= range mo.p {
            wg.Add(1)
            go runProc(ctx context.Context,wg,myCancel <-chProc chan *Data)
    }
    
    
    LOOP:
    for i,_:= range dataSlc {
        
        for k,_:= range mo.p {
            dataSlc[i]->chProc
               
        }
        
         func (){
             switch{
                case resData<-chRes:
                    resDataSlc=append(resDataSlc,resData)
                case <-ctx.Done():
                    return
             }
           
         }()
            func (){
            wg.Wait()
            close(chProc)
            close(chRes)
            }()
    }
            
     if ctx.Err()==nil{
        mo.w.Write(resDataSlc)
     }
    
    resDataSlc = resDataSlc[:0]
    
    }
}

func runProc(ctx *context.Context,wg *sync.WaitGroup, myCancel func(), chProc <-chan *Data,chRes chan *Data)([]*Data, error){
    defer wg.Done()
    var val *Data
                select {
                    case val<-chProc:
                    data,err:=mo.p.Process(val)
                    if err!=nil{
                        myCancel()
                    }
                    data->chRes
                    case <-ctx.Done():
                        return
                }
}