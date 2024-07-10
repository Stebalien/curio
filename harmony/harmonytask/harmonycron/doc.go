/*
	  Workflow:
		HarmonyCron is a regular harmonytask.TaskInterface
		It supports: At(when time.Time, taskType string, sqlTable, sqlRowID, sqlTaskIDColumnName)
			which will add a task to the task engine at the specified time and associate it with the specified row.

		Operation:
			The cron-task will be picked up by Cron runners, which try to avoid being sealers or provers.
			The cron-task is held until the specified time in seconds, then it completes after starting
			the task in the task engine.
*/
package harmonycron
