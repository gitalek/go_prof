package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := in
	for _, stage := range stages {
		out = stageMid(out, done, stage)
	}
	return out
}

func stageMid(in In, done In, stage Stage) Out {
	out := make(Bi)
	go func() {
		defer func() {
			close(out)
		}()
		for v := range in {
			select {
			case <-done:
				return
			default:
				out <- v
			}
		}
	}()
	return stage(out)
}
