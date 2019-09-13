package syntax

type Statement interface {
	GetType() string
	IsCondition() bool
	IsLoop() bool
	IsTrueLoop() bool
	IsWhileLoop() bool
	IsBreak() bool
	IsContinue() bool
	IsProcCall() bool
	IsVarDecl() bool
	IsExpression() bool
	IsReturn() bool
	IsProcReturn() bool
	IsFuncReturn() bool
}

const (
	StatementCondition  = "condition"
	StatementTrueLoop   = "true_loop"
	StatementWhileLoop  = "while_loop"
	StatementBreak      = "break"
	StatementContinue   = "continue"
	StatementProcCall   = "proc_call"
	StatementVarDecl    = "var_decl"
	StatementExpression = "expression"
	StatementProcReturn = "proc_return"
	StatementFuncReturn = "func_return"
)
