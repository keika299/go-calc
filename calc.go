// Copyright (c) 2017 keika299
//
// This software is released under the MIT License.
// http://opensource.org/licenses/mit-license.php

// Package calc exec string expression.
//
// This package only can use operator +, -, *, and /.
// Cannot use round bracket and other operators.
package calc

import (
	"errors"
	"regexp"
	"strconv"
)

var trimSpace = regexp.MustCompile(`[ ]`)
var checkRegExp = regexp.MustCompile(`^[+\-]?\d+(?:\.\d+)?(?:[+\-*/]\d+(?:\.\d+)?)*$`)
var singleOperandRegExp = regexp.MustCompile(`([+\-*/])?(\d+(?:\.\d+)?)`)
var checkConditionalRexExp = regexp.MustCompile(`^([+\-]?\d+(?:\.\d+)?(?:[+\-*/]\d+(?:\.\d+)?)*)(<|>|=|<=|>=)([+\-]?\d+(?:\.\d+)?(?:[+\-*/]\d+(?:\.\d+)?)*)$`)

type expression struct {
}

type block struct {
	operator string
	value    float64
}

// Run return execute expression result.
// Return value type is float64.
func Run(expression string) (float64, error) {

	expression = trimAllSpace(expression)

	if !checkRegExp.MatchString(expression) {
		return 0.0, errors.New("not match expression")
	}

	root, errBuild := build(expression)
	if errBuild != nil {
		return 0.0, errBuild
	}

	return resolve(root)
}

// RunInt return execute expression result.
// Return value type is int.
func RunInt(expression string) (int, error) {
	result, err := Run(expression)

	if err != nil {
		return 0, err
	}

	return int(result), nil
}

func build(expression string) ([]*block, error) {
	singleOperandArray := singleOperandRegExp.FindAllStringSubmatch(expression, -1)

	blocks := make([]*block, 0, len(singleOperandArray))

	for i, single := range singleOperandArray {
		block := new(block)
		if i == 0 {
			block.operator = "+"
			value, _ := strconv.ParseFloat(single[2], 64)
			if single[1] == "-" {
				value = -value
			}
			block.value = value
		} else {
			block.operator = single[1]
			value, _ := strconv.ParseFloat(single[2], 64)
			block.value = value
		}
		blocks = append(blocks, block)
	}

	return blocks, nil
}

func resolve(blocks []*block) (float64, error) {

	if len(blocks) < 2 {
		return blocks[0].value, nil
	}

	operatorList := [][]string{
		{"*", "/"},
		{"+", "-"},
	}

	for _, operators := range operatorList {
		result, err := resolveTargetGroup(blocks, operators)
		if err != nil {
			return 0.0, err
		}
		blocks = result
	}

	if len(blocks) != 1 {
		return 0.0, errors.New("cannot resolve it. please check expression")
	}

	return blocks[0].value, nil
}

func resolveTargetGroup(blocks []*block, target []string) ([]*block, error) {
	for {
		blockLength := len(blocks)

		for i := 1; i < len(blocks); i++ {
			calcBlock, err := checkTarget(blocks[i-1], blocks[i], target)
			if err != nil {
				return nil, err
			}

			if calcBlock != nil {
				blocks = append(append(blocks[:(i-1)], calcBlock), blocks[(i+1):]...)
				break
			}
		}

		if blockLength == len(blocks) {
			return blocks, nil
		}
	}
}

func checkTarget(block1 *block, block2 *block, target []string) (*block, error) {
	for _, operator := range target {
		if block2.operator == operator {
			block := new(block)
			block.operator = block1.operator
			value, err := calc(block1.value, operator, block2.value)
			if err != nil {
				return nil, err
			}
			block.value = value
			return block, nil
		}
	}
	return nil, nil
}

func calc(operand1 float64, operator string, operand2 float64) (float64, error) {
	switch operator {
	case "+":
		return operand1 + operand2, nil
	case "-":
		return operand1 - operand2, nil
	case "*":
		return operand1 * operand2, nil
	case "/":
		if operand2 == 0.0 {
			return 0.0, errors.New("Cannot divide with 0")
		}
		return operand1 / operand2, nil
	}

	return 0.0, errors.New("Invalid operator")
}

func trimAllSpace(s string) string {
	return trimSpace.ReplaceAllString(s, "")
}

// ConditionalExpression returns about it is collect condition.
// When input invalid condition, return false and error object.
func ConditionalExpression(expression string) (bool, error) {
	expression = trimAllSpace(expression)
	targetArray := checkConditionalRexExp.FindAllStringSubmatch(expression, -1)
	if targetArray == nil {
		return false, errors.New("Invalid experssion")
	}

	target := targetArray[0]

	value1, _ := Run(target[1])
	value2, _ := Run(target[3])

	switch target[2] {
	case "=":
		return value1 == value2, nil
	case "<":
		return value1 < value2, nil
	case ">":
		return value1 > value2, nil
	case "<=":
		return value1 <= value2, nil
	case ">=":
		return value1 >= value2, nil
	}

	return false, errors.New("only allow =, <, >, <=,and >=")
}
