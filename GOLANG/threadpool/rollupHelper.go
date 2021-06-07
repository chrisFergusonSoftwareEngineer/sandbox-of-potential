package threadpool

type QueryForRollup struct {
	Id     string
	Date   string
	Result chan int
}

func (query *QueryForRollup) Do() {
	search := SearchTask{
		Date:       query.Date,
		Id:         query.Id,
		ResultChan: make(chan int, 1),
	}

	// dispatch search
	if GetWorkerPool().Enqueue(&search) {

		processSearch := ProcessSearchResult{
			Query:  *query,
			Search: search,
		}
		GetWorkerPool().Enqueue(&processSearch)
	}
}

type ProcessSearchResult struct {
	Query  QueryForRollup
	Search SearchTask
}

func (processSearch *ProcessSearchResult) Do() {
	result, foundResult := <-processSearch.Search.ResultChan

	if foundResult {
		processSearch.Query.Result <- result
	} else {
		close(processSearch.Query.Result)
	}
}
